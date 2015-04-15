/*****************************************************************************
DESCRIPTION Skeleton of the read-driver
*****************************************************************************/
/*-------------------- I n c l u d e F i l e s -------------------------*/
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/fs.h>
#include <linux/uaccess.h>
/*-------------------- C o n s t a n t s ---------------------------------*/
#define DEV_MAJOR 0 // Means to use a dynamic
#define DEV_NAME "simp_rw"
#define SUCCESS 0
#define BUF_SIZE 200 //size for my buffer
/*-------------------- V a r i a b l e s ---------------------------------*/
static int major_number;
static char theBuf[BUF_SIZE]= {0}; // buffer to contain messages
static  loff_t index; //keep track on position between writes
/*-------------------- F u n c t i o n s ---------------------------------*/
static int dev_open(struct inode *inode, struct file *file)
{
    printk("dev_opened. ");
    return 0;
}
static int dev_release(struct inode *inode, struct file *file)
{
    printk("simplkm_read: skeleton device driver closed\n");
    return 0;
}
// read from device
static ssize_t dev_read(struct file *file, char *buf, size_t count, loff_t *ppos)
{
    int msg_size= sizeof(theBuf); // buffer capasity
    if( msg_size <=*ppos)return 0; // trying to read outside buffer area.
    if( msg_size <= *ppos + count)count = msg_size - *ppos; //reading outside message length
    if( copy_to_user(buf,theBuf, count)) { //copy to user space.
        printk(KERN_ERR "Error: copy_to_user.Size of count: %zu. \n",count);
        return -EFAULT;
    }
    memset(&theBuf[0], '\0', sizeof(theBuf)); //message has been read. clear the "inbox".
    index=0; // reset index.
    *ppos+=count;
    return count;
}
// write to device
static ssize_t dev_write (struct file *file, const char __user *buf, size_t count, loff_t *ppos) {
    if ((size_t)BUF_SIZE<= count) { // msg to long. Return
        printk(KERN_WARNING "Message to long. Size count: %zu | capasity: %d",count,BUF_SIZE);
        return -1;
    }
    if ((size_t)BUF_SIZE<index) return -1;  // index outside buffer. Return to prevent overflow
    if ((size_t)BUF_SIZE<index+count) count = BUF_SIZE - index-1;  // adjust count to prevent overflow
    sprintf(theBuf+index,"\n"); // format to add new line for each message.
    printk(KERN_DEBUG "Before Copy: %lld | Count: %zu Strlen(TheBuf): %d",index,count,strlen(theBuf));
    if (copy_from_user(theBuf+index+1, buf, count)) {  // write from userspace (*buf) to kernel space (theBuf)
        printk(KERN_ERR "Error: copy_from_user. Size of count: %zu.",count);
        return -EFAULT;
    }
    printk(KERN_DEBUG "The buffer (theBuf) from kernel space is: %s \n",theBuf);
    *ppos+= count;
    index+=*ppos; // increase index.
    return count;
}
struct file_operations dev_fops= {
    .owner=THIS_MODULE,
    .read=dev_read,
    .write=dev_write,
    .open=dev_open,
    .release=dev_release
};
static int __init dev_init_module(void)
{
    major_number = register_chrdev(0, DEV_NAME, &dev_fops);
    if (major_number < 0)
    {
        printk(KERN_ALERT "Registering char device failed with %d\n", major_number);
        return major_number;
    }
    printk(KERN_INFO "'mknod /dev/%s c %d 0'.\n", DEV_NAME, major_number);
    return SUCCESS;
}
static void __exit dev_cleanup_module(void)
{
    unregister_chrdev(major_number, DEV_NAME);
    printk("simplkm_rw: dev_cleanup_module Device Driver Removed\n");
}
module_init(dev_init_module);
module_exit(dev_cleanup_module);
MODULE_AUTHOR("Morten Mossige, University of Stavanger");
MODULE_DESCRIPTION("Sample Linux device driver");
MODULE_LICENSE("GPL");
