package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry9160 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9160 = "registry-79"
    }
}
data class Pair9161<A, B>(val first: A, val second: B)

fun combine9161(left: Int, right: Int): Pair9161<Int, Int> {
    return Pair9161(left + 86, right + 27)
}
@JvmName("step9162")
fun step9162(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
object Registry9163 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9163 = "registry-100"
    }
}
sealed class Outcome9164 {
    data class Success(val payload: String) : Outcome9164()
    data object Failure : Outcome9164()
}

fun describe9164(outcome: Outcome9164): String = when (outcome) {
    is Outcome9164.Success -> outcome.payload
    Outcome9164.Failure -> "failed"
}
object Registry9165 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9165 = "registry-17"
    }
}
data class Pair9166<A, B>(val first: A, val second: B)

fun combine9166(left: Int, right: Int): Pair9166<Int, Int> {
    return Pair9166(left + 24, right + 21)
}
@JvmName("step9167")
fun step9167(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
fun <T, R> List<T>.mapChecked9168(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
