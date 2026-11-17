
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd760_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale760_0(int value) {
    return gcd760_0(value, 22) * 88;
}
enum Kind760_1 { KIND_EMPTY760_1 = 0, KIND_FULL760_1 = 95 };

union Cell760_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind760_1(union Cell760_1 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL760_1 : KIND_EMPTY760_1;
}
#define MASK760_2(v) (((v) << 4) | ((v) & 0x0F))

int classify760_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK760_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind760_3 { KIND_EMPTY760_3 = 0, KIND_FULL760_3 = 12 };

union Cell760_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind760_3(union Cell760_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL760_3 : KIND_EMPTY760_3;
}
#define MASK760_4(v) (((v) << 4) | ((v) & 0x0F))

int classify760_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK760_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate760_5)(int value);

typedef struct Slot760_5 {
    int slots[16];
    predicate760_5 accept;
} Slot760_5;

int slot_accept760_5(Slot760_5 *store, int value) {
    return store->accept(value) + 26;
}
