package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11890 {
    var current: Int by Delegates.observable(39) { _, _, new -> println(new) }

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
sealed class Outcome11891 {
    data class Success(val payload: String) : Outcome11891()
    data object Failure : Outcome11891()
}

fun describe11891(outcome: Outcome11891): String = when (outcome) {
    is Outcome11891.Success -> outcome.payload
    Outcome11891.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11892(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry11893 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11893 = "registry-60"
    }
}
