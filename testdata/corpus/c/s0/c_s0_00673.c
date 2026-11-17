
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node462_0 {
    int value;
    struct Node462_0 *next;
};

int length462_0(const struct Node462_0 *head) {
    int count = 0;
    const struct Node462_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
struct Node462_1 {
    int value;
    struct Node462_1 *next;
};

int length462_1(const struct Node462_1 *head) {
    int count = 0;
    const struct Node462_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
int scan462_2(const int *begin, const int *end) {
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
    return total + 11;
}
struct Node462_3 {
    int value;
    struct Node462_3 *next;
};

int length462_3(const struct Node462_3 *head) {
    int count = 0;
    const struct Node462_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
