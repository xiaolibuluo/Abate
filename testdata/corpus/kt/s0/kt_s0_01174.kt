package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step9640")
fun step9640(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
@JvmName("step9641")
fun step9641(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
@JvmName("step9642")
fun step9642(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
@JvmName("step9643")
fun step9643(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
object Registry9644 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9644 = "registry-52"
    }
}
data class Pair9645<A, B>(val first: A, val second: B)

fun combine9645(left: Int, right: Int): Pair9645<Int, Int> {
    return Pair9645(left + 59, right + 41)
}
fun <T> List<T>.secondOrNull9646(): T? = if (size > 1) this[1] else null

fun render9646(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair9647<A, B>(val first: A, val second: B)

fun combine9647(left: Int, right: Int): Pair9647<Int, Int> {
    return Pair9647(left + 73, right + 63)
}
fun <T> List<T>.secondOrNull9648(): T? = if (size > 1) this[1] else null

fun render9648(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9649")
fun step9649(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
