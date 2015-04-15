
/*****************************************************************************
DESCRIPTION Skeleton of the read-driver
*****************************************************************************/
/*-------------------- I n c l u d e F i l e s -------------------------*/
#include <linux/module.h> // included for all kernel modules
#include <linux/kernel.h> // included for KERN_INFO
#include <linux/init.h> // included for __init and __exit macros
#include <linux/fs.h>
#include <linux/uaccess.h>
/*-------------------- C o n s t a n t s ---------------------------------*/
#define DEV_MAJOR 0 // Means to use a dynamic
#define DEV_NAME "simp_read"
#define SUCCESS 0
/*-------------------- V a r i a b l e s ---------------------------------*/
static int major_number,n_accessed=0;
static char theBuf[100]={0};
/*-------------------- F u n c t i o n s ---------------------------------*/
// open function - called when the "file" /dev/simp_read is opened in user-space
static int dev_open(struct inode *inode, struct file *file)
{
    n_accessed++; // file has been accessed. Increase number of reads.
    printk("simplkm_read: skeleton device driver open\n");
    return 0;
}
// close function - called when the "file" /dev/simp_read is closed in user-space
static int dev_release(struct inode *inode, struct file *file)
{
    printk("simplkm_read: skeleton device driver closed\n");
    return 0;
}
// read function called when from /dev/simp_read is read
static ssize_t dev_read(struct file *file, char *buf, size_t count, loff_t *ppos)
{
    int bufLen=0;
    sprintf(theBuf, "The Number of driver read= %d\n",n_accessed); //write msg into buffer with formatting
    bufLen = strlen( theBuf);
    printk("simplkm_read: %s times. Buflen=%d\n",theBuf, bufLen );
     if( bufLen <=*ppos)return 0; // trying to read outside buffer area.
	 if( bufLen < *ppos + count)count = bufLen - *ppos; //reading outside message length    
     if( copy_to_user(buf,theBuf, count)){ //copy to user space. 
        printk(KERN_ERR "Error: copy_to_user.Size of count: %zu. \n",count); //
		return -EFAULT;
     } 
	*ppos+=count;  
     return count;
}
// define which file operations are supported
struct file_operations dev_fops =
{
    .owner=THIS_MODULE,
    .read=dev_read,
    .write=NULL,
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
    printk("simplkm_read: dev_cleanup_module Device Driver Removed\n");
}
module_init(dev_init_module);
module_exit(dev_cleanup_module);
MODULE_AUTHOR("Morten Mossige, University of Stavanger");
MODULE_DESCRIPTION("Sample Linux device driver");
MODULE_LICENSE("GPL");
