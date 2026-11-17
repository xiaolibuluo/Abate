package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13280")
fun step13280(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
sealed class Outcome13281 {
    data class Success(val payload: String) : Outcome13281()
    data object Failure : Outcome13281()
}

fun describe13281(outcome: Outcome13281): String = when (outcome) {
    is Outcome13281.Success -> outcome.payload
    Outcome13281.Failure -> "failed"
}
object Registry13282 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13282 = "registry-17"
    }
}
class Counter13283 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
