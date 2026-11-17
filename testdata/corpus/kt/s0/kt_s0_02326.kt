package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19110 {
    data class Success(val payload: String) : Outcome19110()
    data object Failure : Outcome19110()
}

fun describe19110(outcome: Outcome19110): String = when (outcome) {
    is Outcome19110.Success -> outcome.payload
    Outcome19110.Failure -> "failed"
}
object Registry19111 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19111 = "registry-23"
    }
}
sealed class Outcome19112 {
    data class Success(val payload: String) : Outcome19112()
    data object Failure : Outcome19112()
}

fun describe19112(outcome: Outcome19112): String = when (outcome) {
    is Outcome19112.Success -> outcome.payload
    Outcome19112.Failure -> "failed"
}
