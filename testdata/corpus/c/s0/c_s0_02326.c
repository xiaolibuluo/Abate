
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1570_0 {
    int value;
    struct Node1570_0 *next;
};

int length1570_0(const struct Node1570_0 *head) {
    int count = 0;
    const struct Node1570_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
struct Node1570_1 {
    int value;
    struct Node1570_1 *next;
};

int length1570_1(const struct Node1570_1 *head) {
    int count = 0;
    const struct Node1570_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
int scan1570_2(const int *begin, const int *end) {
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
    return total + 59;
}
