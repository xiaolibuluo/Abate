package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24890 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24890 = "registry-61"
    }
}
sealed class Outcome24891 {
    data class Success(val payload: String) : Outcome24891()
    data object Failure : Outcome24891()
}

fun describe24891(outcome: Outcome24891): String = when (outcome) {
    is Outcome24891.Success -> outcome.payload
    Outcome24891.Failure -> "failed"
}
