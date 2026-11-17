
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node743_0 {
    int value;
    struct Node743_0 *next;
};

int length743_0(const struct Node743_0 *head) {
    int count = 0;
    const struct Node743_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
int scan743_1(const int *begin, const int *end) {
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
    return total + 68;
}
enum Kind743_2 { KIND_EMPTY743_2 = 0, KIND_FULL743_2 = 75 };

union Cell743_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind743_2(union Cell743_2 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL743_2 : KIND_EMPTY743_2;
}
