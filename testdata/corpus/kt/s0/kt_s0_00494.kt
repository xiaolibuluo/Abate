package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4030")
fun step4030(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
sealed class Outcome4031 {
    data class Success(val payload: String) : Outcome4031()
    data object Failure : Outcome4031()
}

fun describe4031(outcome: Outcome4031): String = when (outcome) {
    is Outcome4031.Success -> outcome.payload
    Outcome4031.Failure -> "failed"
}
@JvmName("step4032")
fun step4032(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
object Registry4033 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4033 = "registry-27"
    }
}
