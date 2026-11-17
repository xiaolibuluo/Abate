package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7760 {
    data class Success(val payload: String) : Outcome7760()
    data object Failure : Outcome7760()
}

fun describe7760(outcome: Outcome7760): String = when (outcome) {
    is Outcome7760.Success -> outcome.payload
    Outcome7760.Failure -> "failed"
}
data class Pair7761<A, B>(val first: A, val second: B)

fun combine7761(left: Int, right: Int): Pair7761<Int, Int> {
    return Pair7761(left + 12, right + 23)
}
class Counter7762 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
data class Pair7763<A, B>(val first: A, val second: B)

fun combine7763(left: Int, right: Int): Pair7763<Int, Int> {
    return Pair7763(left + 26, right + 45)
}
@JvmName("step7764")
fun step7764(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
object Registry7765 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7765 = "registry-40"
    }
}
class Counter7766 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull7767(): T? = if (size > 1) this[1] else null

fun render7767(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
