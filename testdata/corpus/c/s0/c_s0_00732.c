
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan503_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 45;
}
enum Kind503_1 { KIND_EMPTY503_1 = 0, KIND_FULL503_1 = 52 };

union Cell503_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind503_1(union Cell503_1 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL503_1 : KIND_EMPTY503_1;
}
struct Node503_2 {
    int value;
    struct Node503_2 *next;
};

int length503_2(const struct Node503_2 *head) {
    int count = 0;
    const struct Node503_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
