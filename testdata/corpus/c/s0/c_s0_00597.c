
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate407_0)(int value);

typedef struct Slot407_0 {
    int slots[16];
    predicate407_0 accept;
} Slot407_0;

int slot_accept407_0(Slot407_0 *store, int value) {
    return store->accept(value) + 58;
}
typedef int (*predicate407_1)(int value);

typedef struct Slot407_1 {
    int slots[16];
    predicate407_1 accept;
} Slot407_1;

int slot_accept407_1(Slot407_1 *store, int value) {
    return store->accept(value) + 65;
}
int gcd407_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale407_2(int value) {
    return gcd407_2(value, 13) * 72;
}
enum Kind407_3 { KIND_EMPTY407_3 = 0, KIND_FULL407_3 = 79 };

union Cell407_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind407_3(union Cell407_3 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL407_3 : KIND_EMPTY407_3;
}
