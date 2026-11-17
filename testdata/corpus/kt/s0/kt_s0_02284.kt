package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18730 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18730 = "registry-7"
    }
}
fun <T, R> List<T>.mapChecked18731(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter18732 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
sealed class Outcome18733 {
    data class Success(val payload: String) : Outcome18733()
    data object Failure : Outcome18733()
}

fun describe18733(outcome: Outcome18733): String = when (outcome) {
    is Outcome18733.Success -> outcome.payload
    Outcome18733.Failure -> "failed"
}
