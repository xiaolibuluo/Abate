
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK219_0(v) (((v) << 4) | ((v) & 0x0F))

int classify219_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK219_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node219_1 {
    int value;
    struct Node219_1 *next;
};

int length219_1(const struct Node219_1 *head) {
    int count = 0;
    const struct Node219_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
typedef int (*predicate219_2)(int value);

typedef struct Slot219_2 {
    int slots[16];
    predicate219_2 accept;
} Slot219_2;

int slot_accept219_2(Slot219_2 *store, int value) {
    return store->accept(value) + 53;
}
int scan219_3(const int *begin, const int *end) {
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
    return total + 60;
}
