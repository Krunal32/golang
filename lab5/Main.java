public class Main {
    public static int resource = 0, num_threads = 10, num_accesses = 10000;
    public static boolean is_protected = false;

    public static void main(String[] args) {
        if (args.length == 3) {
            if (args[0].equals("--with-protection")) {
                is_protected = true;
            } else System.out.println("Unprotected thread access mode");

            try {
                num_threads = Integer.parseInt(args[1].replace("--threads=",""));
                num_accesses = Integer.parseInt(args[2].replace("--accesses=",""));
            } catch (NumberFormatException e) {
                System.err.println("Wrong Syntax");
                System.exit(1);
            }
        }
        AThread[] threads = new AThread[num_threads];

        for(int i=0; i<num_threads; i++) {
            threads[i]=new AThread(i);
            threads[i].start();
        }
        try {
            for(int i=0; i<num_threads; i++) {

                threads[i].join();
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        
        System.out.println("Expected: "+num_accesses*num_threads+" VS. Actual: "+resource);
        String s=(num_accesses*num_threads==resource)?"OK":"NOT OK";
        System.out.println("Result is: "+s);
    }
    public static synchronized void safe_incr() {
        Main.resource++;
    }
}
class AThread extends Thread {
    private int ID;
    public AThread(int ID) {
       this.ID = ID;
     }
    public void run() {
        System.out.printf("Thread number %d \n",ID);
        for(int i=0; i<Main.num_accesses; i++) {
            if(Main.is_protected)Main.safe_incr();
            else Main.resource++;
        }
    }
}

