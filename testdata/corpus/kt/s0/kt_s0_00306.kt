package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry2550 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2550 = "registry-22"
    }
}
sealed class Outcome2551 {
    data class Success(val payload: String) : Outcome2551()
    data object Failure : Outcome2551()
}

fun describe2551(outcome: Outcome2551): String = when (outcome) {
    is Outcome2551.Success -> outcome.payload
    Outcome2551.Failure -> "failed"
}
@JvmName("step2552")
fun step2552(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
