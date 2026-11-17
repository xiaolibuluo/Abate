package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12260 {
    data class Success(val payload: String) : Outcome12260()
    data object Failure : Outcome12260()
}

fun describe12260(outcome: Outcome12260): String = when (outcome) {
    is Outcome12260.Success -> outcome.payload
    Outcome12260.Failure -> "failed"
}
data class Pair12261<A, B>(val first: A, val second: B)

fun combine12261(left: Int, right: Int): Pair12261<Int, Int> {
    return Pair12261(left + 42, right + 62)
}
object Registry12262 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12262 = "registry-49"
    }
}
