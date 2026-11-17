package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12040 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12040 = "registry-40"
    }
}
fun <T> List<T>.secondOrNull12041(): T? = if (size > 1) this[1] else null

fun render12041(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12042")
fun step12042(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
sealed class Outcome12043 {
    data class Success(val payload: String) : Outcome12043()
    data object Failure : Outcome12043()
}

fun describe12043(outcome: Outcome12043): String = when (outcome) {
    is Outcome12043.Success -> outcome.payload
    Outcome12043.Failure -> "failed"
}
data class Pair12044<A, B>(val first: A, val second: B)

fun combine12044(left: Int, right: Int): Pair12044<Int, Int> {
    return Pair12044(left + 68, right + 63)
}
class Counter12045 {
    var current: Int by Delegates.observable(75) { _, _, new -> println(new) }

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
object Registry12046 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12046 = "registry-82"
    }
}
