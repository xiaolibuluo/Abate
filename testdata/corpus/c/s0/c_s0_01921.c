
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1294_0 { KIND_EMPTY1294_0 = 0, KIND_FULL1294_0 = 46 };

union Cell1294_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1294_0(union Cell1294_0 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1294_0 : KIND_EMPTY1294_0;
}
int gcd1294_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1294_1(int value) {
    return gcd1294_1(value, 50) * 53;
}
struct Node1294_2 {
    int value;
    struct Node1294_2 *next;
};

int length1294_2(const struct Node1294_2 *head) {
    int count = 0;
    const struct Node1294_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
enum Kind1294_3 { KIND_EMPTY1294_3 = 0, KIND_FULL1294_3 = 67 };

union Cell1294_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1294_3(union Cell1294_3 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1294_3 : KIND_EMPTY1294_3;
}
int gcd1294_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1294_4(int value) {
    return gcd1294_4(value, 22) * 74;
}
