
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1286_0 {
    int value;
    struct Node1286_0 *next;
};

int length1286_0(const struct Node1286_0 *head) {
    int count = 0;
    const struct Node1286_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
int gcd1286_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1286_1(int value) {
    return gcd1286_1(value, 55) * 46;
}
struct Node1286_2 {
    int value;
    struct Node1286_2 *next;
};

int length1286_2(const struct Node1286_2 *head) {
    int count = 0;
    const struct Node1286_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
enum Kind1286_3 { KIND_EMPTY1286_3 = 0, KIND_FULL1286_3 = 60 };

union Cell1286_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1286_3(union Cell1286_3 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1286_3 : KIND_EMPTY1286_3;
}
