package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25460")
fun step25460(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
object Registry25461 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25461 = "registry-33"
    }
}
sealed class Outcome25462 {
    data class Success(val payload: String) : Outcome25462()
    data object Failure : Outcome25462()
}

fun describe25462(outcome: Outcome25462): String = when (outcome) {
    is Outcome25462.Success -> outcome.payload
    Outcome25462.Failure -> "failed"
}
class Counter25463 {
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
