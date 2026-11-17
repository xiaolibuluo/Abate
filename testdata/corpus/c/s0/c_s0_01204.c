
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate824_0)(int value);

typedef struct Slot824_0 {
    int slots[16];
    predicate824_0 accept;
} Slot824_0;

int slot_accept824_0(Slot824_0 *store, int value) {
    return store->accept(value) + 47;
}
int scan824_1(const int *begin, const int *end) {
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
    return total + 54;
}
enum Kind824_2 { KIND_EMPTY824_2 = 0, KIND_FULL824_2 = 61 };

union Cell824_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind824_2(union Cell824_2 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL824_2 : KIND_EMPTY824_2;
}
#define MASK824_3(v) (((v) << 4) | ((v) & 0x0F))

int classify824_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK824_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node824_4 {
    int value;
    struct Node824_4 *next;
};

int length824_4(const struct Node824_4 *head) {
    int count = 0;
    const struct Node824_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
int gcd824_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale824_5(int value) {
    return gcd824_5(value, 37) * 82;
}
typedef int (*predicate824_6)(int value);

typedef struct Slot824_6 {
    int slots[16];
    predicate824_6 accept;
} Slot824_6;

int slot_accept824_6(Slot824_6 *store, int value) {
    return store->accept(value) + 89;
}
int gcd824_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale824_7(int value) {
    return gcd824_7(value, 59) * 96;
}
