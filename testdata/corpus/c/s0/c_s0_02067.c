
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1395_0 {
    int value;
    struct Node1395_0 *next;
};

int length1395_0(const struct Node1395_0 *head) {
    int count = 0;
    const struct Node1395_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
int scan1395_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 8;
}
int gcd1395_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1395_2(int value) {
    return gcd1395_2(value, 36) * 15;
}
enum Kind1395_3 { KIND_EMPTY1395_3 = 0, KIND_FULL1395_3 = 22 };

union Cell1395_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1395_3(union Cell1395_3 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1395_3 : KIND_EMPTY1395_3;
}
struct Node1395_4 {
    int value;
    struct Node1395_4 *next;
};

int length1395_4(const struct Node1395_4 *head) {
    int count = 0;
    const struct Node1395_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
