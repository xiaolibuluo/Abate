
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate772_0)(int value);

typedef struct Slot772_0 {
    int slots[16];
    predicate772_0 accept;
} Slot772_0;

int slot_accept772_0(Slot772_0 *store, int value) {
    return store->accept(value) + 50;
}
enum Kind772_1 { KIND_EMPTY772_1 = 0, KIND_FULL772_1 = 57 };

union Cell772_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind772_1(union Cell772_1 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL772_1 : KIND_EMPTY772_1;
}
int gcd772_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale772_2(int value) {
    return gcd772_2(value, 67) * 64;
}
struct Node772_3 {
    int value;
    struct Node772_3 *next;
};

int length772_3(const struct Node772_3 *head) {
    int count = 0;
    const struct Node772_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
#define MASK772_4(v) (((v) << 4) | ((v) & 0x0F))

int classify772_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK772_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd772_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale772_5(int value) {
    return gcd772_5(value, 39) * 85;
}
#define MASK772_6(v) (((v) << 4) | ((v) & 0x0F))

int classify772_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK772_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
