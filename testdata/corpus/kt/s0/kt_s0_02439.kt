package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20070 {
    data class Success(val payload: String) : Outcome20070()
    data object Failure : Outcome20070()
}

fun describe20070(outcome: Outcome20070): String = when (outcome) {
    is Outcome20070.Success -> outcome.payload
    Outcome20070.Failure -> "failed"
}
data class Pair20071<A, B>(val first: A, val second: B)

fun combine20071(left: Int, right: Int): Pair20071<Int, Int> {
    return Pair20071(left + 10, right + 39)
}
object Registry20072 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20072 = "registry-17"
    }
}
