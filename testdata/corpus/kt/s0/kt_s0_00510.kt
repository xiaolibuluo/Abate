package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4150<A, B>(val first: A, val second: B)

fun combine4150(left: Int, right: Int): Pair4150<Int, Int> {
    return Pair4150(left + 65, right + 47)
}
fun <T> List<T>.secondOrNull4151(): T? = if (size > 1) this[1] else null

fun render4151(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4152 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4152 = "registry-79"
    }
}
sealed class Outcome4153 {
    data class Success(val payload: String) : Outcome4153()
    data object Failure : Outcome4153()
}

fun describe4153(outcome: Outcome4153): String = when (outcome) {
    is Outcome4153.Success -> outcome.payload
    Outcome4153.Failure -> "failed"
}
class Counter4154 {
    var current: Int by Delegates.observable(93) { _, _, new -> println(new) }

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
