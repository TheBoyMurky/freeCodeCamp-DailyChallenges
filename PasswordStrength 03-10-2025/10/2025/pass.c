#include <stdio.h>
#include <string.h>


int main(void) {

    char pass_input[32];
    int pass_len;
    int pass_up_letters_flag = 0;
    int pass_low_letters_flag = 0;
    int pass_nums_flag = 0;
    int pass_specials_flag = 0;
    int pass_strength = 0;

    puts("Test you password strength.");
    puts("Rules to follow:");
    puts("It is at least 8 characters long.");
    puts("It contains both uppercase and lowercase letters.");
    puts("Password input:");
    fgets(pass_input, 32, stdin);

    pass_len = strlen(pass_input);
    if (pass_len >= 8) {
        pass_strength++;
    }

    for (int i; i <= pass_len; i++) {
        // Upper
        if (pass_input[i] >= 65 && pass_input[i] <= 90) {
            pass_up_letters_flag=1;
        }
        // Lower
        if (pass_input[i] >= 97 && pass_input[i] <= 122) {
            pass_low_letters_flag=1;
        }
        // Nums
        if (pass_input[i] >= 48 && pass_input[i] <= 57) {
            pass_nums_flag=1;
        }
        // Special chars
        switch (pass_input[i]) {
            // !
            case 33:
                pass_specials_flag=1;
                break;
            // #
            case 35:
                pass_specials_flag=1;
                break;
            // $
            case 36:
                pass_specials_flag=1;
                break;
            // %
            case 37:
                pass_specials_flag=1;
                break;
            // &
            case 38:
                pass_specials_flag=1;
                break;
            // *
            case 42:
                pass_specials_flag=1;
                break;
            // @
            case 64:
                pass_specials_flag=1;
                break;
            // ^
            case 94:
                pass_specials_flag=1;
                break;
        }
    }
    
    // Sum all flags
    if (pass_up_letters_flag && pass_low_letters_flag)
        pass_strength++;
    if (pass_nums_flag)
        pass_strength++;
    if (pass_specials_flag)
        pass_strength++;

    // Final print
    switch (pass_strength) {
        case 1:
            puts("weak...");
            break;
        case 2:
        case 3:
            puts("medium.");
            break;
        case 4:
            puts("Strong!");
            break;
        default:
            puts("WEAKLING!");
    }
        
    return 0;
}