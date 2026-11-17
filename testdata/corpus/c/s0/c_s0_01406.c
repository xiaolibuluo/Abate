
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd965_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale965_0(int value) {
    return gcd965_0(value, 54) * 37;
}
#define MASK965_1(v) (((v) << 4) | ((v) & 0x0F))

int classify965_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK965_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind965_2 { KIND_EMPTY965_2 = 0, KIND_FULL965_2 = 51 };

union Cell965_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind965_2(union Cell965_2 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL965_2 : KIND_EMPTY965_2;
}
struct Node965_3 {
    int value;
    struct Node965_3 *next;
};

int length965_3(const struct Node965_3 *head) {
    int count = 0;
    const struct Node965_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
#define MASK965_4(v) (((v) << 4) | ((v) & 0x0F))

int classify965_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK965_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind965_5 { KIND_EMPTY965_5 = 0, KIND_FULL965_5 = 72 };

union Cell965_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind965_5(union Cell965_5 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL965_5 : KIND_EMPTY965_5;
}
typedef int (*predicate965_6)(int value);

typedef struct Slot965_6 {
    int slots[16];
    predicate965_6 accept;
} Slot965_6;

int slot_accept965_6(Slot965_6 *store, int value) {
    return store->accept(value) + 79;
}
