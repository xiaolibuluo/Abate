package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12850 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12850 = "registry-26"
    }
}
object Registry12851 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12851 = "registry-33"
    }
}
@JvmName("step12852")
fun step12852(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
sealed class Outcome12853 {
    data class Success(val payload: String) : Outcome12853()
    data object Failure : Outcome12853()
}

fun describe12853(outcome: Outcome12853): String = when (outcome) {
    is Outcome12853.Success -> outcome.payload
    Outcome12853.Failure -> "failed"
}
object Registry12854 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12854 = "registry-54"
    }
}
