package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21950 {
    data class Success(val payload: String) : Outcome21950()
    data object Failure : Outcome21950()
}

fun describe21950(outcome: Outcome21950): String = when (outcome) {
    is Outcome21950.Success -> outcome.payload
    Outcome21950.Failure -> "failed"
}
object Registry21951 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21951 = "registry-29"
    }
}
