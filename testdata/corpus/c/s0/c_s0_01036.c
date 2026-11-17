
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK710_0(v) (((v) << 4) | ((v) & 0x0F))

int classify710_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK710_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK710_1(v) (((v) << 4) | ((v) & 0x0F))

int classify710_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK710_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate710_2)(int value);

typedef struct Slot710_2 {
    int slots[16];
    predicate710_2 accept;
} Slot710_2;

int slot_accept710_2(Slot710_2 *store, int value) {
    return store->accept(value) + 34;
}
struct Node710_3 {
    int value;
    struct Node710_3 *next;
};

int length710_3(const struct Node710_3 *head) {
    int count = 0;
    const struct Node710_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
typedef int (*predicate710_4)(int value);

typedef struct Slot710_4 {
    int slots[16];
    predicate710_4 accept;
} Slot710_4;

int slot_accept710_4(Slot710_4 *store, int value) {
    return store->accept(value) + 48;
}
int gcd710_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale710_5(int value) {
    return gcd710_5(value, 32) * 55;
}
