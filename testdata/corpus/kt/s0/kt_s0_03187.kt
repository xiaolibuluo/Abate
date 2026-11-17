package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair26190<A, B>(val first: A, val second: B)

fun combine26190(left: Int, right: Int): Pair26190<Int, Int> {
    return Pair26190(left + 5, right + 42)
}
data class Pair26191<A, B>(val first: A, val second: B)

fun combine26191(left: Int, right: Int): Pair26191<Int, Int> {
    return Pair26191(left + 12, right + 53)
}
object Registry26192 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26192 = "registry-19"
    }
}
@JvmName("step26193")
fun step26193(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
object Registry26194 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26194 = "registry-33"
    }
}
object Registry26195 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26195 = "registry-40"
    }
}
data class Pair26196<A, B>(val first: A, val second: B)

fun combine26196(left: Int, right: Int): Pair26196<Int, Int> {
    return Pair26196(left + 47, right + 47)
}
fun <T> List<T>.secondOrNull26197(): T? = if (size > 1) this[1] else null

fun render26197(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair26198<A, B>(val first: A, val second: B)

fun combine26198(left: Int, right: Int): Pair26198<Int, Int> {
    return Pair26198(left + 61, right + 69)
}
