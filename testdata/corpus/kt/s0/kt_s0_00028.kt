package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome230 {
    data class Success(val payload: String) : Outcome230()
    data object Failure : Outcome230()
}

fun describe230(outcome: Outcome230): String = when (outcome) {
    is Outcome230.Success -> outcome.payload
    Outcome230.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked231(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry232 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG232 = "registry-27"
    }
}
class Counter233 {
    var current: Int by Delegates.observable(34) { _, _, new -> println(new) }

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
