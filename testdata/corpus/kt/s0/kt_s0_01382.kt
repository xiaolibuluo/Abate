package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome11380 {
    data class Success(val payload: String) : Outcome11380()
    data object Failure : Outcome11380()
}

fun describe11380(outcome: Outcome11380): String = when (outcome) {
    is Outcome11380.Success -> outcome.payload
    Outcome11380.Failure -> "failed"
}
object Registry11381 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11381 = "registry-62"
    }
}
object Registry11382 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11382 = "registry-69"
    }
}
object Registry11383 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11383 = "registry-76"
    }
}
