
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1746_0 {
    int value;
    struct Node1746_0 *next;
};

int length1746_0(const struct Node1746_0 *head) {
    int count = 0;
    const struct Node1746_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
struct Node1746_1 {
    int value;
    struct Node1746_1 *next;
};

int length1746_1(const struct Node1746_1 *head) {
    int count = 0;
    const struct Node1746_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
#define MASK1746_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1746_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1746_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1746_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1746_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1746_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1746_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1746_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1746_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1746_5)(int value);

typedef struct Slot1746_5 {
    int slots[16];
    predicate1746_5 accept;
} Slot1746_5;

int slot_accept1746_5(Slot1746_5 *store, int value) {
    return store->accept(value) + 40;
}
int scan1746_6(const int *begin, const int *end) {
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
    return total + 47;
}
