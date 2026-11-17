package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17670 {
    data class Success(val payload: String) : Outcome17670()
    data object Failure : Outcome17670()
}

fun describe17670(outcome: Outcome17670): String = when (outcome) {
    is Outcome17670.Success -> outcome.payload
    Outcome17670.Failure -> "failed"
}
object Registry17671 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17671 = "registry-91"
    }
}
fun <T, R> List<T>.mapChecked17672(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter17673 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
