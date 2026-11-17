package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18920")
fun step18920(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
fun <T> List<T>.secondOrNull18921(): T? = if (size > 1) this[1] else null

fun render18921(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18922<A, B>(val first: A, val second: B)

fun combine18922(left: Int, right: Int): Pair18922<Int, Int> {
    return Pair18922(left + 74, right + 38)
}
object Registry18923 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18923 = "registry-81"
    }
}
class Counter18924 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull18925(): T? = if (size > 1) this[1] else null

fun render18925(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome18926 {
    data class Success(val payload: String) : Outcome18926()
    data object Failure : Outcome18926()
}

fun describe18926(outcome: Outcome18926): String = when (outcome) {
    is Outcome18926.Success -> outcome.payload
    Outcome18926.Failure -> "failed"
}
object Registry18927 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18927 = "registry-12"
    }
}
