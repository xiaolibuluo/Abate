package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step23070")
fun step23070(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
@JvmName("step23071")
fun step23071(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
fun <T> List<T>.secondOrNull23072(): T? = if (size > 1) this[1] else null

fun render23072(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step23073")
fun step23073(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
fun <T, R> List<T>.mapChecked23074(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
