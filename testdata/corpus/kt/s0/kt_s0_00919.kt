package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter7530 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

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
object Registry7531 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7531 = "registry-101"
    }
}
fun <T> List<T>.secondOrNull7532(): T? = if (size > 1) this[1] else null

fun render7532(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7533<A, B>(val first: A, val second: B)

fun combine7533(left: Int, right: Int): Pair7533<Int, Int> {
    return Pair7533(left + 18, right + 67)
}
sealed class Outcome7534 {
    data class Success(val payload: String) : Outcome7534()
    data object Failure : Outcome7534()
}

fun describe7534(outcome: Outcome7534): String = when (outcome) {
    is Outcome7534.Success -> outcome.payload
    Outcome7534.Failure -> "failed"
}
