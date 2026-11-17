package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry19750 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19750 = "registry-72"
    }
}
data class Pair19751<A, B>(val first: A, val second: B)

fun combine19751(left: Int, right: Int): Pair19751<Int, Int> {
    return Pair19751(left + 79, right + 59)
}
sealed class Outcome19752 {
    data class Success(val payload: String) : Outcome19752()
    data object Failure : Outcome19752()
}

fun describe19752(outcome: Outcome19752): String = when (outcome) {
    is Outcome19752.Success -> outcome.payload
    Outcome19752.Failure -> "failed"
}
fun <T> List<T>.secondOrNull19753(): T? = if (size > 1) this[1] else null

fun render19753(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter19754 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
