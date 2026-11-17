package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17890(): T? = if (size > 1) this[1] else null

fun render17890(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step17891")
fun step17891(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
fun <T> List<T>.secondOrNull17892(): T? = if (size > 1) this[1] else null

fun render17892(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair17893<A, B>(val first: A, val second: B)

fun combine17893(left: Int, right: Int): Pair17893<Int, Int> {
    return Pair17893(left + 100, right + 60)
}
object Registry17894 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17894 = "registry-10"
    }
}
data class Pair17895<A, B>(val first: A, val second: B)

fun combine17895(left: Int, right: Int): Pair17895<Int, Int> {
    return Pair17895(left + 17, right + 21)
}
object Registry17896 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17896 = "registry-24"
    }
}
@JvmName("step17897")
fun step17897(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
fun <T> List<T>.secondOrNull17898(): T? = if (size > 1) this[1] else null

fun render17898(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
