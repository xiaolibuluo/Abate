package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6530(): T? = if (size > 1) this[1] else null

fun render6530(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6531<A, B>(val first: A, val second: B)

fun combine6531(left: Int, right: Int): Pair6531<Int, Int> {
    return Pair6531(left + 62, right + 16)
}
data class Pair6532<A, B>(val first: A, val second: B)

fun combine6532(left: Int, right: Int): Pair6532<Int, Int> {
    return Pair6532(left + 69, right + 27)
}
@JvmName("step6533")
fun step6533(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
@JvmName("step6534")
fun step6534(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
@JvmName("step6535")
fun step6535(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
