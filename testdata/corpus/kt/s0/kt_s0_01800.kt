package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14790<A, B>(val first: A, val second: B)

fun combine14790(left: Int, right: Int): Pair14790<Int, Int> {
    return Pair14790(left + 26, right + 53)
}
object Registry14791 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14791 = "registry-33"
    }
}
class Counter14792 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

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
sealed class Outcome14793 {
    data class Success(val payload: String) : Outcome14793()
    data object Failure : Outcome14793()
}

fun describe14793(outcome: Outcome14793): String = when (outcome) {
    is Outcome14793.Success -> outcome.payload
    Outcome14793.Failure -> "failed"
}
