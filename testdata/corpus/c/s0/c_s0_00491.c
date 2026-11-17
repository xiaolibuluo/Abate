
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate338_0)(int value);

typedef struct Slot338_0 {
    int slots[16];
    predicate338_0 accept;
} Slot338_0;

int slot_accept338_0(Slot338_0 *store, int value) {
    return store->accept(value) + 34;
}
int gcd338_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale338_1(int value) {
    return gcd338_1(value, 68) * 41;
}
struct Node338_2 {
    int value;
    struct Node338_2 *next;
};

int length338_2(const struct Node338_2 *head) {
    int count = 0;
    const struct Node338_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
#define MASK338_3(v) (((v) << 4) | ((v) & 0x0F))

int classify338_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK338_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate338_4)(int value);

typedef struct Slot338_4 {
    int slots[16];
    predicate338_4 accept;
} Slot338_4;

int slot_accept338_4(Slot338_4 *store, int value) {
    return store->accept(value) + 62;
}
int scan338_5(const int *begin, const int *end) {
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
    return total + 69;
}
typedef int (*predicate338_6)(int value);

typedef struct Slot338_6 {
    int slots[16];
    predicate338_6 accept;
} Slot338_6;

int slot_accept338_6(Slot338_6 *store, int value) {
    return store->accept(value) + 76;
}
enum Kind338_7 { KIND_EMPTY338_7 = 0, KIND_FULL338_7 = 83 };

union Cell338_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind338_7(union Cell338_7 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL338_7 : KIND_EMPTY338_7;
}
