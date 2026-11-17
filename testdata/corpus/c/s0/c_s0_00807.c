
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd558_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale558_0(int value) {
    return gcd558_0(value, 11) * 81;
}
int gcd558_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale558_1(int value) {
    return gcd558_1(value, 22) * 88;
}
int gcd558_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale558_2(int value) {
    return gcd558_2(value, 33) * 95;
}
typedef int (*predicate558_3)(int value);

typedef struct Slot558_3 {
    int slots[16];
    predicate558_3 accept;
} Slot558_3;

int slot_accept558_3(Slot558_3 *store, int value) {
    return store->accept(value) + 5;
}
enum Kind558_4 { KIND_EMPTY558_4 = 0, KIND_FULL558_4 = 12 };

union Cell558_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind558_4(union Cell558_4 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL558_4 : KIND_EMPTY558_4;
}
