package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step17320")
fun step17320(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
data class Pair17321<A, B>(val first: A, val second: B)

fun combine17321(left: Int, right: Int): Pair17321<Int, Int> {
    return Pair17321(left + 24, right + 66)
}
data class Pair17322<A, B>(val first: A, val second: B)

fun combine17322(left: Int, right: Int): Pair17322<Int, Int> {
    return Pair17322(left + 31, right + 16)
}
data class Pair17323<A, B>(val first: A, val second: B)

fun combine17323(left: Int, right: Int): Pair17323<Int, Int> {
    return Pair17323(left + 38, right + 27)
}
object Registry17324 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17324 = "registry-45"
    }
}
fun <T> List<T>.secondOrNull17325(): T? = if (size > 1) this[1] else null

fun render17325(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17326 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
