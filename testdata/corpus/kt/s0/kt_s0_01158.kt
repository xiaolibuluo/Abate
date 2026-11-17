package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9530<A, B>(val first: A, val second: B)

fun combine9530(left: Int, right: Int): Pair9530<Int, Int> {
    return Pair9530(left + 75, right + 31)
}
fun <T> List<T>.secondOrNull9531(): T? = if (size > 1) this[1] else null

fun render9531(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter9532 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
@JvmName("step9533")
fun step9533(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
object Registry9534 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9534 = "registry-6"
    }
}
object Registry9535 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9535 = "registry-13"
    }
}
