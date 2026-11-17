package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6560 {
    data class Success(val payload: String) : Outcome6560()
    data object Failure : Outcome6560()
}

fun describe6560(outcome: Outcome6560): String = when (outcome) {
    is Outcome6560.Success -> outcome.payload
    Outcome6560.Failure -> "failed"
}
data class Pair6561<A, B>(val first: A, val second: B)

fun combine6561(left: Int, right: Int): Pair6561<Int, Int> {
    return Pair6561(left + 101, right + 37)
}
class Counter6562 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
object Registry6563 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6563 = "registry-18"
    }
}
fun <T> List<T>.secondOrNull6564(): T? = if (size > 1) this[1] else null

fun render6564(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
