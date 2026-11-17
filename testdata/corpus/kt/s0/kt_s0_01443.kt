package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11900 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11900 = "registry-52"
    }
}
object Registry11901 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11901 = "registry-59"
    }
}
fun <T, R> List<T>.mapChecked11902(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11903<A, B>(val first: A, val second: B)

fun combine11903(left: Int, right: Int): Pair11903<Int, Int> {
    return Pair11903(left + 73, right + 15)
}
object Registry11904 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11904 = "registry-80"
    }
}
@JvmName("step11905")
fun step11905(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
