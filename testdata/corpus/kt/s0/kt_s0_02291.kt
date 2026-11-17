package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18800 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
data class Pair18801<A, B>(val first: A, val second: B)

fun combine18801(left: Int, right: Int): Pair18801<Int, Int> {
    return Pair18801(left + 8, right + 65)
}
object Registry18802 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18802 = "registry-15"
    }
}
fun <T> List<T>.secondOrNull18803(): T? = if (size > 1) this[1] else null

fun render18803(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18804")
fun step18804(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
object Registry18805 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18805 = "registry-36"
    }
}
fun <T> List<T>.secondOrNull18806(): T? = if (size > 1) this[1] else null

fun render18806(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
