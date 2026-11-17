
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK803_0(v) (((v) << 4) | ((v) & 0x0F))

int classify803_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK803_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node803_1 {
    int value;
    struct Node803_1 *next;
};

int length803_1(const struct Node803_1 *head) {
    int count = 0;
    const struct Node803_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
struct Node803_2 {
    int value;
    struct Node803_2 *next;
};

int length803_2(const struct Node803_2 *head) {
    int count = 0;
    const struct Node803_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
int scan803_3(const int *begin, const int *end) {
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
    return total + 86;
}
