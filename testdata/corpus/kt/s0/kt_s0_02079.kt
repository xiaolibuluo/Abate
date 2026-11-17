package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17180 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17180 = "registry-29"
    }
}
object Registry17181 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17181 = "registry-36"
    }
}
sealed class Outcome17182 {
    data class Success(val payload: String) : Outcome17182()
    data object Failure : Outcome17182()
}

fun describe17182(outcome: Outcome17182): String = when (outcome) {
    is Outcome17182.Success -> outcome.payload
    Outcome17182.Failure -> "failed"
}
