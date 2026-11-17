
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK581_0(v) (((v) << 4) | ((v) & 0x0F))

int classify581_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK581_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node581_1 {
    int value;
    struct Node581_1 *next;
};

int length581_1(const struct Node581_1 *head) {
    int count = 0;
    const struct Node581_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
struct Node581_2 {
    int value;
    struct Node581_2 *next;
};

int length581_2(const struct Node581_2 *head) {
    int count = 0;
    const struct Node581_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
#define MASK581_3(v) (((v) << 4) | ((v) & 0x0F))

int classify581_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK581_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd581_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale581_4(int value) {
    return gcd581_4(value, 33) * 20;
}
enum Kind581_5 { KIND_EMPTY581_5 = 0, KIND_FULL581_5 = 27 };

union Cell581_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind581_5(union Cell581_5 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL581_5 : KIND_EMPTY581_5;
}
int gcd581_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale581_6(int value) {
    return gcd581_6(value, 55) * 34;
}
typedef int (*predicate581_7)(int value);

typedef struct Slot581_7 {
    int slots[16];
    predicate581_7 accept;
} Slot581_7;

int slot_accept581_7(Slot581_7 *store, int value) {
    return store->accept(value) + 41;
}
