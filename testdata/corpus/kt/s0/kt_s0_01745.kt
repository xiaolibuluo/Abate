package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14300 {
    data class Success(val payload: String) : Outcome14300()
    data object Failure : Outcome14300()
}

fun describe14300(outcome: Outcome14300): String = when (outcome) {
    is Outcome14300.Success -> outcome.payload
    Outcome14300.Failure -> "failed"
}
@JvmName("step14301")
fun step14301(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
sealed class Outcome14302 {
    data class Success(val payload: String) : Outcome14302()
    data object Failure : Outcome14302()
}

fun describe14302(outcome: Outcome14302): String = when (outcome) {
    is Outcome14302.Success -> outcome.payload
    Outcome14302.Failure -> "failed"
}
object Registry14303 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14303 = "registry-89"
    }
}
