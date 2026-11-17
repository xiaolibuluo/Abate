
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate778_0)(int value);

typedef struct Slot778_0 {
    int slots[16];
    predicate778_0 accept;
} Slot778_0;

int slot_accept778_0(Slot778_0 *store, int value) {
    return store->accept(value) + 31;
}
typedef int (*predicate778_1)(int value);

typedef struct Slot778_1 {
    int slots[16];
    predicate778_1 accept;
} Slot778_1;

int slot_accept778_1(Slot778_1 *store, int value) {
    return store->accept(value) + 38;
}
enum Kind778_2 { KIND_EMPTY778_2 = 0, KIND_FULL778_2 = 45 };

union Cell778_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind778_2(union Cell778_2 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL778_2 : KIND_EMPTY778_2;
}
#define MASK778_3(v) (((v) << 4) | ((v) & 0x0F))

int classify778_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK778_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK778_4(v) (((v) << 4) | ((v) & 0x0F))

int classify778_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK778_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node778_5 {
    int value;
    struct Node778_5 *next;
};

int length778_5(const struct Node778_5 *head) {
    int count = 0;
    const struct Node778_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
int gcd778_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale778_6(int value) {
    return gcd778_6(value, 31) * 73;
}
struct Node778_7 {
    int value;
    struct Node778_7 *next;
};

int length778_7(const struct Node778_7 *head) {
    int count = 0;
    const struct Node778_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
