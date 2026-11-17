
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node771_0 {
    int value;
    struct Node771_0 *next;
};

int length771_0(const struct Node771_0 *head) {
    int count = 0;
    const struct Node771_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
#define MASK771_1(v) (((v) << 4) | ((v) & 0x0F))

int classify771_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK771_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate771_2)(int value);

typedef struct Slot771_2 {
    int slots[16];
    predicate771_2 accept;
} Slot771_2;

int slot_accept771_2(Slot771_2 *store, int value) {
    return store->accept(value) + 51;
}
typedef int (*predicate771_3)(int value);

typedef struct Slot771_3 {
    int slots[16];
    predicate771_3 accept;
} Slot771_3;

int slot_accept771_3(Slot771_3 *store, int value) {
    return store->accept(value) + 58;
}
int gcd771_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale771_4(int value) {
    return gcd771_4(value, 21) * 65;
}
int gcd771_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale771_5(int value) {
    return gcd771_5(value, 32) * 72;
}
enum Kind771_6 { KIND_EMPTY771_6 = 0, KIND_FULL771_6 = 79 };

union Cell771_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind771_6(union Cell771_6 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL771_6 : KIND_EMPTY771_6;
}
