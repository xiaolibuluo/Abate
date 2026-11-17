package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10990 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10990 = "registry-33"
    }
}
object Registry10991 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10991 = "registry-40"
    }
}
object Registry10992 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10992 = "registry-47"
    }
}
data class Pair10993<A, B>(val first: A, val second: B)

fun combine10993(left: Int, right: Int): Pair10993<Int, Int> {
    return Pair10993(left + 54, right + 49)
}
@JvmName("step10994")
fun step10994(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
sealed class Outcome10995 {
    data class Success(val payload: String) : Outcome10995()
    data object Failure : Outcome10995()
}

fun describe10995(outcome: Outcome10995): String = when (outcome) {
    is Outcome10995.Success -> outcome.payload
    Outcome10995.Failure -> "failed"
}
