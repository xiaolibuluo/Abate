
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK434_0(v) (((v) << 4) | ((v) & 0x0F))

int classify434_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK434_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node434_1 {
    int value;
    struct Node434_1 *next;
};

int length434_1(const struct Node434_1 *head) {
    int count = 0;
    const struct Node434_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
typedef int (*predicate434_2)(int value);

typedef struct Slot434_2 {
    int slots[16];
    predicate434_2 accept;
} Slot434_2;

int slot_accept434_2(Slot434_2 *store, int value) {
    return store->accept(value) + 35;
}
enum Kind434_3 { KIND_EMPTY434_3 = 0, KIND_FULL434_3 = 42 };

union Cell434_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind434_3(union Cell434_3 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL434_3 : KIND_EMPTY434_3;
}
int scan434_4(const int *begin, const int *end) {
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
    return total + 49;
}
typedef int (*predicate434_5)(int value);

typedef struct Slot434_5 {
    int slots[16];
    predicate434_5 accept;
} Slot434_5;

int slot_accept434_5(Slot434_5 *store, int value) {
    return store->accept(value) + 56;
}
#define MASK434_6(v) (((v) << 4) | ((v) & 0x0F))

int classify434_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK434_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node434_7 {
    int value;
    struct Node434_7 *next;
};

int length434_7(const struct Node434_7 *head) {
    int count = 0;
    const struct Node434_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}
